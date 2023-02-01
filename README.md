# godomain


    super.admin.gmail.google.com
    \_______________/ \____/ \_/
          record      domain tld


Usage:

    d := godomain.Parse("super.admin.gmail.google.com")
	fmt.Println(d)          // print: super.admin.gmail.google.com
	fmt.Println(d.Tld())    // print: com
	fmt.Println(d.Domain()) // print: google.com
	fmt.Println(d.Record()) // print: super.admin.gmail
