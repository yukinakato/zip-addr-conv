#!/usr/bin/ruby

# run "gem install nokogiri" before running this script

require "nokogiri"
require "open-uri"
require "net/http"

doc = Nokogiri::HTML(URI.open("https://www.post.japanpost.jp/zipcode/dl/kogaki-zip.html"))
retrieved = doc.xpath('//*[@id="main-box"]/div/div[1]/p/small').text.to_s

last = FileTest.exist?("lastupdate.txt") ? File.read("lastupdate.txt") : ""

if last != retrieved
  File.open("lastupdate.txt", "wb") do |f|
    f.write(retrieved)
  end
  http = Net::HTTP.new("notify-api.line.me", port = 443)
  http.use_ssl = true
  token = ENV["LN_TOKEN"]
  message = retrieved.gsub(/更新/, "") + "に更新されました。"
  http.post('/api/notify', "message=#{message}", header = {Authorization: "Bearer #{token}"})
end
