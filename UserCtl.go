package aipsetup

import (
	"errors"
	"fmt"
	"os"
	"path"
	"sort"

	"github.com/AnimusPEXUS/shadowusers"
	"github.com/AnimusPEXUS/utils/filetools"
	"github.com/AnimusPEXUS/utils/logger"
)

const SYS_UID_MAX = 999

var USERS = map[int]string{

	// TODO: this list requires serious cleanup

	// # users for groups

	// # lspecial users 1-9
	1: "nobody",
	2: "nogroup",
	3: "bin",
	4: "ftp",
	5: "mail",
	6: "adm",
	7: "gdm",
	8: "wheel",

	// # terminals 10-19
	10: "pts",
	11: "tty",

	// # devices 20-38
	20: "disk",
	21: "usb",
	22: "flash",
	23: "mouse",
	24: "lp",
	25: "floppy",
	26: "video",
	27: "audio",
	28: "cdrom",
	29: "tape",
	30: "pulse",
	31: "pulse-access",
	32: "usbfs",
	33: "usbdev",
	34: "usbbus",
	35: "usblist",
	36: "alsa",

	// # daemons 40-99
	39: "colord",

	40: "messagebus",
	41: "sshd",
	42: "haldaemon",
	//	43: "clamav",
	44: "mysql",
	45: "exim",
	46: "postgres",
	47: "httpd",
	48: "cron",
	//	49: "mrim",
	//	50: "icq",
	//	51: "pyvkt",
	//	52: "j2j",
	//	53: "gnunet",
	//	54: "ejabberd",
	55: "cupsd",
	//	56: "bandersnatch",
	//	57: "torrent",
	58: "ssl",
	//	59: "dovecot",
	//	60: "dovenull",
	//	61: "spamassassin",
	//	62: "yacy",
	//	63: "irc",
	//	64: "hub",
	//	65: "cynin",
	//	66: "mailman",
	//	67: "asterisk",
	//	68: "bitcoin",
	//	69: "adch",

	//	70: "dialout",
	71: "kmem",
	72: "polkituser",
	//	73: "nexuiz",
	//	74: "couchdb",
	75: "polkitd",
	76: "kvm",

	90: "mine", // TODO: remember what it is. minetest?

	91: "utmp",
	92: "lock",
	93: "avahi",
	94: "avahi-autoipd",
	95: "netdev",
	//	96: "freenet",
	//	97: "jabberd2",
	//	98: "mongodb",
	99: "aipsetupserv",

	100: "systemd-bus-proxy",
	101: "systemd-network",
	102: "systemd-resolve",
	103: "systemd-timesync",
	104: "systemd-journal",
	105: "systemd-journal-gateway",
	106: "systemd-journal-remote",
	107: "systemd-journal-upload",

	200: "tor",
	//	201: "shinken",
}

type UserCtl struct {
	sys *System
}

func NewUserCtl(sys *System) *UserCtl {
	self := new(UserCtl)
	self.sys = sys
	return self
}

func (self *UserCtl) UserKeysSortedSlice() []int {
	ret := make([]int, 0)

	for k, _ := range USERS {
		ret = append(ret, k)
	}
	sort.Ints(ret)
	return ret
}

func (self *UserCtl) UserIdByName(name string) (int, error) {
	for k, v := range USERS {
		if v == name {
			return k, nil
		}
	}
	return -1, errors.New("not found")
}

func (self *UserCtl) UserNameById(id int) (string, error) {
	name, ok := USERS[id]
	if !ok {
		return "", errors.New("not found")
	}
	return name, nil
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

	for k, v := range USERS {
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

	ks := self.UserKeysSortedSlice()

	for _, k := range ks {

		i := USERS[k]

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

func (self *UserCtl) ResetSystemPermissions(log *logger.Logger) error {

	was_errors := false

	loginfo := func(txt string) {
		if log != nil {
			log.Info(txt)
		}
	}

	logerr := func(err error) {
		was_errors = true
		if log != nil {
			log.Error(err)
		}
	}

	r := self.sys.Root()
	{
		loginfo("chown root: /")
		err := os.Chown(r, 0, 0)
		if err != nil {
			logerr(err)
		}

		loginfo("chmod 755 /")
		err = os.Chmod(r, 0755)
		if err != nil {
			logerr(err)
		}
	}

	{
		p := path.Join(r, "/tmp")

		loginfo("chown root: /tmp")
		err := os.Chown(p, 0, 0)
		if err != nil {
			logerr(err)
		}

		loginfo("chmod 1777 /tmp")
		err = os.Chmod(p, 01777)
		if err != nil {
			logerr(err)
		}
	}

	// NOTE: imo Lilith system doesn't need this
	//chown root:mail /var/mail
	//chmod 1777 /var/mail

	//# polkit settings
	//chown root:root /etc/polkit-1/localauthority
	//chmod 0700 /etc/polkit-1/localauthority
	{
		p := path.Join(r, "/etc", "polkit-1", "localauthority")

		loginfo("chown root: " + p)
		err := os.Chown(p, 0, 0)
		if err != nil {
			logerr(err)
		}

		loginfo("chmod 0700: " + p)
		err = os.Chmod(p, 0700)
		if err != nil {
			logerr(err)
		}
	}

	//#chown root:root /var/lib/polkit-1
	//#chmod 0700 /var/lib/polkit-1

	//chown root:root /etc/pam.d/polkit-1
	//chmod 0700 /etc/pam.d/polkit-1
	{
		p := path.Join(r, "/etc", "pam.d", "polkit-1")

		loginfo("chown root: " + p)
		err := os.Chown(p, 0, 0)
		if err != nil {
			logerr(err)
		}

		loginfo("chmod 0700: " + p)
		err = os.Chmod(p, 0700)
		if err != nil {
			logerr(err)
		}
	}

	//# systemd service files

	//for i in \
	//    '/usr/lib/systemd/system' \
	//    '/usr/lib/systemd/user' \
	//    '/etc/systemd/system' \
	//    '/etc/systemd/user'
	//do

	//    chmod 0755 "$i"
	//    find "$i" -type d -exec chmod 755 '{}' ';'
	//    find "$i" -type f -exec chmod 644 '{}' ';'

	//done

	{
		for _, i := range []string{
			"/usr/lib/systemd/system",
			"/usr/lib/systemd/user",
			"/etc/systemd/system",
			"/etc/systemd/user",
		} {
			loginfo("systemd config files: " + i)
			err := filetools.Walk(
				path.Join(r, i),
				func(pth string, dirs, files []os.FileInfo) error {
					for _, i := range dirs {
						fp := path.Join(pth, path.Base(i.Name()))
						err := os.Chown(fp, 0, 0)
						if err != nil {
							return err
						}
						err = os.Chmod(fp, 0755)
						if err != nil {
							return err
						}
					}
					for _, i := range files {
						fp := path.Join(pth, path.Base(i.Name()))
						err := os.Chown(fp, 0, 0)
						if err != nil {
							return err
						}
						err = os.Chmod(fp, 0644)
						if err != nil {
							return err
						}
					}
					return nil
				},
			)
			if err != nil {
				logerr(err)
			}
		}
	}

	//chmod 4755 /usr/libexec/dbus-daemon-launch-helper
	//chmod 4755 /usr/lib/polkit-1/polkit-agent-helper-1
	//chmod 4755 /usr/bin/pkexec
	{
		for _, i := range []string{
			"/usr/libexec/dbus-daemon-launch-helper",
			"/usr/lib/polkit-1/polkit-agent-helper-1",
			"/usr/bin/pkexec",
		} {

			p := path.Join(r, i)
			loginfo("chown root: " + p)
			err := os.Chown(p, 0, 0)
			if err != nil {
				logerr(err)
			}

			loginfo("chmod 4755 " + p)
			err = os.Chmod(p, 04755)
			if err != nil {
				logerr(err)
			}
		}
	}

	//chmod 4755 "`which su`"
	//chmod 4755 "`which sudo`"
	{
		for _, i := range []string{
			"su",
			"sudo",
		} {
			loginfo("which " + i + "?")
			pth, err := filetools.Which(i, []string{})
			if err != nil {
				logerr(err)
			} else {
				loginfo("chown root: " + pth)
				err := os.Chown(pth, 0, 0)
				if err != nil {
					logerr(err)
				}

				loginfo("chmod 4755 " + pth)
				err = os.Chmod(pth, 04755)
				if err != nil {
					logerr(err)
				}
			}
		}
	}

	//# chmod 4755 "`which mount`"
	//# chmod 4755 "`which weston-launch`"
	if was_errors {
		return errors.New("was errors")
	}

	return nil
}
