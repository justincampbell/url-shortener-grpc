require 'bundler/setup'
$LOAD_PATH.unshift File.dirname(__FILE__)

require 'shortener_pb'
require 'shortener_services_pb'

HOST = ENV.fetch('SHORTENER_HOST', 'localhost')
PORT = ENV.fetch('SHORTENER_PORT', 1901)

def handle(command, arg)
  stub = ShortenerService::Stub.new("#{HOST}:#{PORT}", :this_channel_is_insecure)

  case command
  when 'shorten'
    request = Shorten::Request.new(url: arg)
    response = stub.shorten(request)
    puts response.token
  when 'expand'
    request = Expand::Request.new(token: arg)
    response = stub.expand(request)
    puts response.url
  when 'exit'
    exit 0
  end
rescue GRPC::InvalidArgument, GRPC::NotFound, GRPC::Unavailable => e
  puts e.details
end

loop do
  puts 'shorten [url] | expand [token] | exit'
  input = gets
  command, arg = input.split(' ')

  handle command, arg

  puts
end
