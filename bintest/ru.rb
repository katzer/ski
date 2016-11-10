
vkergh

def vkergh
  puts " jahada " unless ENV.fetch('jahada')
rescue SyntaxError
  puts "hasd"
end
