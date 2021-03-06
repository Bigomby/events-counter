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
	"strings"

	rdkafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/redBorder/rbforwarder"
)

// PrintVersion displays the application version.
func PrintVersion() {
	_, s := rdkafka.LibraryVersion()
	prettyPubKey := strings.Replace(PubKey, `\n`, "\n", -1)

	fmt.Printf("Events Counter :: %s\n", version)
	fmt.Printf("rbforwarder    :: %s\n", rbforwarder.Version)
	fmt.Printf("Go             :: %s\n", runtime.Version())
	fmt.Printf("librdkafka     :: %s\n", s)
	fmt.Println()
	fmt.Println(prettyPubKey)
}
