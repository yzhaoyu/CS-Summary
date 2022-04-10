## Golang中"omitempty"关键字学习

#### 用法
在Golang中我们经常在JSON数据与struct数据类型之间进行转换；通过在struct中使用tag可以指定JSON的key，例如在表示一个地址数据的时候，JSON数据如下所示：
```json
{
    "planet": "Earth",
    "continent": "Asia",
    "country": " China",
    "province": "Guangdong"
}
```

与之对应的结构体可以表示如下：

```go
type Location struct {
	Star     string `json:"planet"`         // 星球
	Continet string `json:"continent"`      // 洲际
	Nation   string `json:"country"`        // 国家
	Province string `json:"province"`       // 省自治区/直辖市/特别行政区
	City     string `json:"city,omitempty"` // 城市
}
```

如果我们需要忽略`Location`结构体中的某个字段，例如不需要对外提供City字段的值时，可以采用omitempty属性，就可以让输出的json数据不带city字段，不包含该字段的默认值。

#### 陷阱
**陷阱一**
使用 omitempty 也有些小陷阱，一个是该关键字无法忽略掉嵌套结构体。继续用`Location`结构体举例，这回我们想要往地址结构体中加一个新 field 来表示经纬度，如果缺乏相关的数据，暂时可以忽略。新的结构体定义如下所示

```go
type Location struct {
	Star       string       `json:"planet"`
	Continet   string       `json:"continent"`
	Nation     string       `json:"country"`
	Province   string       `json:"province"`
	City       string       `json:"city,omitempty"`
	Coordinate coordinate   `json:"coordinate,omitempty"` 
}

type coordinate struct {
	Lat float64 `json:"latitude"`
	Lng float64 `json:"longitude"`
}
```

读入原来的地址数据，处理后序列化输出，我们就会发现即使加上了 omitempty 关键字，输出的 json 还是带上了一个空的坐标信息

```json
{
    "planet": "Earth",
    "continent": "Asia",
    "nation": "China",
    "province": "Guangdong",
    "coordinate": {
        "latitude": 0,
        "longitude": 0
    }
}
```

为了达到我们想要的效果，可以把坐标定义为指针类型，这样 Golang 就能知道一个指针的“空值”是多少了，否则面对一个我们自定义的结构， Golang 是猜不出我们想要的空值的。于是有了如下的结构体定义

```go
type Location struct {
	Star       string       `json:"planet"`
	Continet   string       `json:"continent"`
	Nation     string       `json:"country"`
	Province   string       `json:"province"`
	City       string       `json:"city,omitempty"`
	Coordinate *coordinate  `json:"coordinate,omitempty"` 
}

type coordinate struct {
	Lat float64 `json:"latitude"`
	Lng float64 `json:"longitude"`
}
```

此时相应输出为：

```json
{
    "planet": "Earth",
    "continent": "Asia",
    "country": " China",
    "province": "Guangdong"
}
```

**陷阱二**
另一个“陷阱”是，对于用 omitempty 定义的 field ，如果给它赋的值恰好等于默认空值的话，在转为 json 之后也不会输出这个 field 。比如说上面定义的经纬度坐标结构体，如果我们将经纬度两个 field 都加上 omitempty

```go
type coordinate struct {
	Lat float64 `json:"latitude"`
	Lng float64 `json:"longitude"`
}
```

然后我们对非洲几内亚湾的“原点坐标”非常感兴趣，于是编写了如下代码

```go
func main() {
	cData := `{
		"latitude": 0.0,
		"longitude": 0.0
	}`
	c := new(coordinate)
	json.Unmarshal([]byte(cData), &c)

        // 具体处理逻辑...

	coordinateBytes, _ := json.MarshalIndent(c, "", "    ")
	fmt.Printf("%s\n", string(coordinateBytes))
}
```

最终我们得到了一个 `{}`

这个坐标消失不见了！但我们的设想是，如果一个地点没有经纬度信息，则悬空，这没有问题，但对于“原点坐标”，我们在确切知道它的经纬度的情况下，（0.0, 0.0）仍然被忽略了。正确的写法也是将结构体内的定义改为指针

```go
type coordinate struct {
	Lat *float64 `json:"latitude"`
	Lng *float64 `json:"longitude"`
}
```

这样空值就从 float64 的 0.0 变为了指针类型的 nil ，我们就能看到正确的经纬度输出。

```json
{
    "latitude": 0,
    "longitude": 0
}
```
