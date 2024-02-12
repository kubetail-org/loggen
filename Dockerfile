# Copyright 2024 Andres Morey
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM golang:1.21.6 AS builder

RUN mkdir app
WORKDIR /app

# install dependencies (for docker cache)
COPY go.mod .
COPY go.sum .
RUN go mod download

# copy code
COPY . .

# build executable
RUN CGO_ENABLED=0 go build -o loggen ./main.go

ENTRYPOINT ["./loggen"]
CMD []

# -----------------------------------------------------------

FROM scratch AS final

# copy app
COPY --from=builder /app/loggen /app/loggen

WORKDIR /app

ENTRYPOINT ["./loggen"]
CMD []
