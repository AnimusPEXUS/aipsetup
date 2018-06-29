package aipsetup

import (
	"fmt"
	"os"
	"path"

	"github.com/AnimusPEXUS/aipsetup/basictypes"
	"github.com/AnimusPEXUS/shadowusers"
)

type UserCtl struct {
	sys *System
}

func NewUserCtl(sys *System) *UserCtl {
	self := new(UserCtl)
	self.sys = sys
	return self
}

func (self *UserCtl) CalcDaemonHomeDir(root string, daemon_name string) string {
	return path.Join(root, basictypes.DIRNAME_DAEMONS, daemon_name)
}

func (self *UserCtl) CalcUserHomeDir(root string, user_name string) (ret string) {

	if user_name == "root" {
		ret = path.Join(root, "root")
	} else {
		ret = path.Join(root, "home", user_name)
	}

	return
}

func (self *UserCtl) RecreateUserDB() error {
	su_ctl := shadowusers.NewCtl(path.Join(self.sys.Root(), "etc"))
	err := su_ctl.ReadAll()
	if err != nil {
		return err
	}

	// TODO: should DB validity be checked? - probably!

	predefined_normal_user_groups := []string{
		"pts", "tty", "pulse-access", "audio", "kvm", "video",
	}

	root_password := ""

	normal_passwd_records := make([]*shadowusers.Passwd, 0)
	normal_shadow_records := make([]*shadowusers.Shadow, 0)

	//	normal_user_uids := set.NewSetInt()

	for _, i := range su_ctl.Passwds.Passwds {
		if i.UserId >= 1000 {
			normal_passwd_records = append(normal_passwd_records, i)
		}
	}

	for _, i := range normal_passwd_records {
		for _, j := range su_ctl.Shadows.Shadows {
			if j.Login == i.Login {
				normal_shadow_records = append(normal_shadow_records, j)
			}
		}
	}

	for _, i := range su_ctl.Shadows.Shadows {
		if i.Login == "root" {
			root_password = i.Password
			break
		}
	}

	su_ctl.NewAll()

	self.makeUser("root", 0, root_password, []string{}, su_ctl)

	for _, i := range normal_passwd_records {
		if i.Login == "root" || i.UserId == 0 {
			continue
		}
		shadow_pass := "!"
		for _, j := range normal_shadow_records {
			if j.Login == i.Login {
				shadow_pass = j.Password
			}
		}
		self.makeUser(
			i.Login,
			i.UserId,
			shadow_pass,
			predefined_normal_user_groups,
			su_ctl,
		)
	}

	su_ctl.Pth = path.Join("/root", "tmp", "newuserdb")
	fmt.Println("new db going to be placed under", su_ctl.Pth)
	os.MkdirAll(su_ctl.Pth, 0700)

	for k, v := range basictypes.USERS {
		self.makeDaemon(v, k, su_ctl)
	}

	su_ctl.WriteAll()
	if err != nil {
		return err
	}

	return nil
}

func (self *UserCtl) makeUser(
	login string,
	id int,
	password string,
	additional_groups []string,
	suctl *shadowusers.Ctl,
) {
	passwd := &shadowusers.Passwd{}
	shadow := &shadowusers.Shadow{}
	group := &shadowusers.Group{}
	gshadow := &shadowusers.GShadow{}

	passwd.Password = "x"
	passwd.Shell = "/bin/bash"

	shadow.AccountExpirationDays = -1
	shadow.InactivityPeriodDays = -1
	shadow.LastChangeDays = -1
	shadow.MaxAgeDays = -1
	shadow.MinAgeDays = -1
	shadow.WarningPeriodDays = -1
	shadow.Password = password

	group.UserList = []string{login}
	group.UserList = append(group.UserList, additional_groups...)
	group.Password = "x"

	gshadow.Name = login
	gshadow.Password = "!!"

	if id == 0 {
		passwd.Login = "root"
		passwd.UserId = 0
		passwd.GroupId = 0
		passwd.Home = "/root"

		shadow.Login = "root"

		group.Name = "root"
		group.GID = 0
	} else {
		passwd.Login = login
		passwd.UserId = id
		passwd.GroupId = id
		passwd.Home = path.Join("/home", login)

		shadow.Login = login

		group.Name = login
		group.GID = id
	}

	suctl.Passwds.Passwds = append(suctl.Passwds.Passwds, passwd)
	suctl.Shadows.Shadows = append(suctl.Shadows.Shadows, shadow)
	suctl.Groups.Groups = append(suctl.Groups.Groups, group)
	suctl.GShadows.GShadows = append(suctl.GShadows.GShadows, gshadow)
}

func (self *UserCtl) makeDaemon(
	login string,
	id int,
	//password string,
	//	additional_groups []string,
	suctl *shadowusers.Ctl,
) {
	passwd := &shadowusers.Passwd{}
	shadow := &shadowusers.Shadow{}
	group := &shadowusers.Group{}
	gshadow := &shadowusers.GShadow{}

	passwd.Login = login
	passwd.Password = "x"
	passwd.UserId = id
	passwd.GroupId = id
	passwd.Home = path.Join("/daemons", login)
	passwd.Shell = "/bin/nologin"

	shadow.AccountExpirationDays = -1
	shadow.InactivityPeriodDays = -1
	shadow.LastChangeDays = -1
	shadow.MaxAgeDays = -1
	shadow.MinAgeDays = -1
	shadow.WarningPeriodDays = -1
	shadow.Password = "!!"
	shadow.Login = login

	group.Name = login
	group.GID = id
	group.Password = "x"
	group.UserList = []string{login}
	//	group.UserList = append(group.UserList, additional_groups...)

	gshadow.Name = login
	gshadow.Password = "!!"

	suctl.Passwds.Passwds = append(suctl.Passwds.Passwds, passwd)
	suctl.Shadows.Shadows = append(suctl.Shadows.Shadows, shadow)
	suctl.Groups.Groups = append(suctl.Groups.Groups, group)
	suctl.GShadows.GShadows = append(suctl.GShadows.GShadows, gshadow)
}

func (self *UserCtl) RecreateDaemonHomes() error {

	ks := basictypes.UserKeysSortedSlice()

	for _, k := range ks {

		i := basictypes.USERS[k]

		daemon_home_path := self.CalcDaemonHomeDir(self.sys.Root(), i)

		_, err := os.Stat(daemon_home_path)
		if err != nil {
			if !os.IsNotExist(err) {
				return err
			}
		}

		err = os.MkdirAll(daemon_home_path, 0700)
		if err != nil {
			return err
		}

		err = os.Chown(daemon_home_path, k, k)
		if err != nil {
			return err
		}

		err = os.Chmod(daemon_home_path, 0700)
		if err != nil {
			return err
		}

		// TODO: port here standard ssl db access related permissions

		// /daemons related excerpt from aipsetup3. (ejabberd is probably no more)

		//chmod 750 /daemons/ejabberd
		//chmod 750 /daemons/ejabberd/var
		//chmod 750 /daemons/ejabberd/var/www
		//chmod -R 750 /daemons/ejabberd/var/www/logs

		//chmod -R 750 /daemons/ssl

	}

	{
		p := path.Join(self.sys.Root(), "/", basictypes.DIRNAME_DAEMONS)
		err := os.Chown(p, 0, 0)
		if err != nil {
			return err
		}

		err = os.Chmod(p, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

// TODO
//func (self *UserCtl) RecreateUserHomes() {
//}
