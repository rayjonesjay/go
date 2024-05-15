echo -e 'Go\nHello\r12ere'
echo -e 'Hello\n\b'
echo -e 'Hel\nlo\b'
echo -e '\nHello\n\n'
echo -e '\nHello World\n\n'
echo -e '\nHello\tWorld\n\n'
echo -e '\\nHello\tWorld\n\n'

echo -e 'Hello\0World'
echo -e 'Hello\\0World'
echo -e 'Hello\0012World'
echo -e 'Hello\\0012World'
echo -e 'Hello\012World'
echo -e 'Hello\\012World'

echo -e '\r'
echo -e '\\r'
echo -e 'Hello\rGood-morning'
echo -e 'Hello\\rGood-morning'

echo -e 'Good-morning\vBarack'

echo -e '\x61\x62'
echo -e '\\x61\\x62'


echo -e '\x61\x62\x0A\x61\x62'
echo -e '\x61\x62\x0A\x61\x62\x0A'

echo -e 't\x61\x62b\x0At\x61\x62b'
echo -e 't\x61\x62b\x0At\x61\x62b\x0A'

echo -e "Hello\f\fWorld\fThere\fab"
echo -e "Hello\f\fWorld\fThere\f"

echo -e "Ring The Bell\a"
echo -e 'hello There 1 to 2!' # expect "hello There 1 to 2!"
echo -e "1a\"#FdwHywR&/()=" # expect "1a"#FdwHywR&/()="

echo -e '\!" #$%&'"'"'()*+,-./' # expect "\!" #$%&'()*+,-./"
