go-beforeeach
=============

Go equivalent of `@Before` and `@After` in Java.
Check example for usage.

Install
=======
```shell
go get -u github.com/jknair0/beforeeach
```

Examples
========
```go
import (
    beforeEach "github.com/jknair0/beforeeach"
    "testing"
)

var it = beforeEach.Create(setUp, tearDown)

func setUp() {
	...
}

func tearDown() {
	...
}

func TestCalculator_Add(t *testing.T) {
    it(func() {
    	// your test goes here
    })
}
```

License
=======
```
   Copyright 2020 JayaKrishnan Nair

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.

```
