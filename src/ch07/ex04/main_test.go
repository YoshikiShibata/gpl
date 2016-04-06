package main

func ExampleParse() {
	contents := `<html>
	<a href="http://yshibata.blog.so-net.ne.jp"/>
	<a href="http://www001.upp.so-net.ne.jp/yshibata" />
	</html>`

	Parse(contents)
	// Output:
	// http://yshibata.blog.so-net.ne.jp
	// http://www001.upp.so-net.ne.jp/yshibata
}
