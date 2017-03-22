// Simple utility for counting messages on a Kafka topic.
//
// Copyright (C) 2017 ENEO Tecnologia SL
// Author: Diego Fernández Barrera <bigomby@gmail.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"runtime"

	rdkafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

// PrintVersion displays the application version.
func PrintVersion() {
	_, s := rdkafka.LibraryVersion()
	fmt.Printf("Events Counter\t:: %s\n", version)
	fmt.Printf("Go\t\t:: %s\n", runtime.Version())
	fmt.Printf("librdkafka\t:: %s\n", s)
}
