# fiberBindingHeader


## _*fiberBindingHeaderService.InitFiberBindingHeader()*_ to bind with default header tag "header"
1. use tag binding with value required to must bind this value

```

type Header struct {
	Version   string  `header:"version" binding:"required"`
	Language  string  `header:"language" binding:"required"`
	Longitude float64 `header:"long"`
	Latitude  float64 `header:"lat"`
	Action    string  `header:"action"`
	I         int     `header:"i"`
	I8        int8    `header:"i8"`
	I16       int16   `header:"i16"`
	I32       int32   `header:"i32"`
	I64       int64   `header:"i64"`
	U         uint    `header:"u"`
	U8        uint8   `header:"u8"`
	U16       uint16  `header:"u16"`
	U32       uint32  `header:"u32"`
	U64       uint64  `header:"u64"`
	B        bool    `header:"bool"`
}

func middleware(c *fiber.Ctx) error {
	var hhh Header
	fbh := fiberBindingHeaderService.InitFiberBindingHeader()
	err := fbh.BindFiberHeader(&hhh, c)
	if err != nil {
		return err
	}
	fmt.Println(hhh.Latitude)
	fmt.Println(hhh.Longitude)
	fmt.Println(hhh.Version)
	fmt.Println(hhh.Language)
	fmt.Println(hhh.I)
	fmt.Println(hhh.I8)
	fmt.Println(hhh.I16)
	fmt.Println(hhh.I32)
	fmt.Println(hhh.I64)
	fmt.Println(hhh.U)
	fmt.Println(hhh.U8)
	fmt.Println(hhh.U16)
	fmt.Println(hhh.U32)
	fmt.Println(hhh.U64)
	fmt.Println(hhh.B)
	return c.Next()
}
```

`------------------------------------------------------------`

## _fiberBindingHeaderService.InitFiberBindingHeaderWithTag("newHeader")_ to bind with custom header tag "newHeader"
```

type Header struct {
	Version   string  `newHeader:"version" binding:"required"`
	Language  string  `newHeader:"language" binding:"required"`
	Longitude float64 `newHeader:"long"`
	Latitude  float64 `newHeader:"lat"`
	Action    string  `newHeader:"action"`
	I         int     `newHeader:"i"`
	I8        int8    `newHeader:"i8"`
	I16       int16   `newHeader:"i16"`
	I32       int32   `newHeader:"i32"`
	I64       int64   `newHeader:"i64"`
	U         uint    `newHeader:"u"`
	U8        uint8   `newHeader:"u8"`
	U16       uint16  `newHeader:"u16"`
	U32       uint32  `newHeader:"u32"`
	U64       uint64  `newHeader:"u64"`
	B         bool    `newHeader:"bool"`
}

func middleware(c *fiber.Ctx) error {
	var hhh Header
	fbh := fiberBindingHeaderService.InitFiberBindingHeaderWithTag("newHeader")
	err := fbh.BindFiberHeader(&hhh, c)
	if err != nil {
		return err
	}
	fmt.Println(hhh.Latitude)
	fmt.Println(hhh.Longitude)
	fmt.Println(hhh.Version)
	fmt.Println(hhh.Language)
	fmt.Println(hhh.I)
	fmt.Println(hhh.I8)
	fmt.Println(hhh.I16)
	fmt.Println(hhh.I32)
	fmt.Println(hhh.I64)
	fmt.Println(hhh.U)
	fmt.Println(hhh.U8)
	fmt.Println(hhh.U16)
	fmt.Println(hhh.U32)
	fmt.Println(hhh.U64)
	fmt.Println(hhh.B)
	return c.Next()
}
```