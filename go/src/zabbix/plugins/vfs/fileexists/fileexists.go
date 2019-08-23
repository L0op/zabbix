/*
** Zabbix
** Copyright (C) 2001-2019 Zabbix SIA
**
** This program is free software; you can redistribute it and/or modify
** it under the terms of the GNU General Public License as published by
** the Free Software Foundation; either version 2 of the License, or
** (at your option) any later version.
**
** This program is distributed in the hope that it will be useful,
** but WITHOUT ANY WARRANTY; without even the implied warranty of
** MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
** GNU General Public License for more details.
**
** You should have received a copy of the GNU General Public License
** along with this program; if not, write to the Free Software
** Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
**/

package fileexists

import (
	"errors"
	"zabbix/internal/plugin"
	"zabbix/pkg/std"
)

// Plugin -
type Plugin struct {
	plugin.Base
}

var impl Plugin

// Export -
func (p *Plugin) Export(key string, params []string, ctx plugin.ContextProvider) (result interface{}, err error) {
	if len(params) != 1 {
		return nil, errors.New("Wrong number of parameters")
	}
	if "" == params[0] {
		return nil, errors.New("Invalid first parameter")
	}
	ret := 0

	if f, err := stdOs.Stat(params[0]); err == nil {
		if mode := f.Mode(); mode.IsRegular() {
			ret = 1
		}
	} else if stdOs.IsExist(err) {
		ret = 1
	}
	return ret, nil

}

var stdOs std.Os

func init() {
	plugin.RegisterMetric(&impl, "existance", "vfs.file.exists", "Returns if file exists or not.")
	stdOs = std.NewOs()
}
