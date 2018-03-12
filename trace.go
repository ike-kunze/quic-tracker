/*
    Maxime Piraux's master's thesis
    Copyright (C) 2017-2018  Maxime Piraux

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Affero General Public License version 3
	as published by the Free Software Foundation.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero General Public License for more details.

    You should have received a copy of the GNU Affero General Public License
    along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package masterthesis

type Trace struct {
	Commit          string                 `json:"commit"`
	Scenario        string                 `json:"scenario"`
	ScenarioVersion int                    `json:"scenario_version"`
	Host            string                 `json:"host"`
	Ip              string                 `json:"ip"`
	Results         map[string]interface{} `json:"results"`
	StartedAt       int64                  `json:"started_at"`
	Duration        uint64                 `json:"duration"`
	ErrorCode       uint8                  `json:"error_code"`
	Stream          []TracePacket          `json:"stream"`
}

type Direction string

const ToServer Direction = "to_server"
const ToClient Direction = "to_client"

type TracePacket struct {
	Direction Direction `json:"direction"`
	Timestamp int64     `json:"timestamp"`
	Data      []byte    `json:"data"`
}
