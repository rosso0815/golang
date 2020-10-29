#!perl -w

use SOAP::Lite;

print SOAP::Lite                                             
-> uri('http://www.soaplite.com/Demo')                                             
-> proxy('http://localhost:1080/pfistera/cgi/perl/soap.cgi')
-> hi()                                                    
-> result;

print "\n";
