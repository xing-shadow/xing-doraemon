/*
@Time : 2020/8/25 11:33
@Author : wangyl
@File : ladap.go
@Software: GoLand
*/
package gobal

import (
	"github.com/go-ldap/ldap"

	"xing-doraemon/pkg/auth/ldaputil"
	"xing-doraemon/pkg/setting/alterGateway"
)

func InitLdap(ldapCfg alterGateway.Ldap) error {
	if ldapCfg.Enabled {
		_ldapCfg := &ldaputil.LdapConfig{
			Url:          ldapCfg.LdapUrl,
			BaseDN:       ldapCfg.LdapBaseDn,
			Scope:        parseLdapScope(ldapCfg.LdapScope),
			BindUsername: ldapCfg.LdapSearchDn,
			BindPassword: ldapCfg.LdapSearchPassword,
			Filter:       ldapCfg.LdapFilter,
		}
		ldaputil.InitLdap(_ldapCfg)
		return nil
	} else {
		return nil
	}
}

func parseLdapScope(scope int) int {
	switch scope {
	case 0:
		return ldap.ScopeBaseObject
	case 1:
		return ldap.ScopeSingleLevel
	case 2:
		return ldap.ScopeWholeSubtree
	default:
		return ldap.ScopeWholeSubtree
	}
}
