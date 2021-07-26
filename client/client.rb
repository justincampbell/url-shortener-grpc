require 'bundler/setup'
require_relative 'shortener_pb'
require_relative 'shortener_services_pb'

def handle(command, arg)
  stub = ShortenerService::Stub.new("localhost:1901", :this_channel_is_insecure)

  case command
  when "shorten"
    request = Shorten::Request.new(url: arg)
    response = stub.shorten(request)
    puts response.token
  when "expand"
    request = Expand::Request.new(token: arg)
    response = stub.expand(request)
    puts response.url
  when "exit"
    exit 0
  end
rescue GRPC::InvalidArgument, GRPC::NotFound => e
  puts e.details
end

loop do
  puts "shorten [url] | expand [token] | exit"
  input = gets
  command, arg = input.split(" ")

  handle command, arg

  puts
end
