# run "gem install rubyzip" before running this script

require "net/http"
require "zip"
require "csv"
require "nkf"

def ConvertWideAndNarrow(str)
  tmp = str
  # wide to narrow
  tmp = tmp.gsub(/０/, "0")
  tmp = tmp.gsub(/１/, "1")
  tmp = tmp.gsub(/２/, "2")
  tmp = tmp.gsub(/３/, "3")
  tmp = tmp.gsub(/４/, "4")
  tmp = tmp.gsub(/５/, "5")
  tmp = tmp.gsub(/６/, "6")
  tmp = tmp.gsub(/７/, "7")
  tmp = tmp.gsub(/８/, "8")
  tmp = tmp.gsub(/９/, "9")
  # narrow to wide
  tmp = tmp.gsub(/\(/, "（")
  tmp = tmp.gsub(/\)/, "）")
  tmp = tmp.gsub(/-/, "－")
  tmp = tmp.gsub(/</, "＜")
  tmp = tmp.gsub(/>/, "＞")
end

client = Net::HTTP.new("www.post.japanpost.jp", 443)
client.use_ssl = true
body = client.get("/zipcode/dl/kogaki/zip/ken_all.zip").body
File.open("ken_all.zip", "wb") do |f|
  f.write(body)
end

Zip::File.open("ken_all.zip") do |zip_file|
  entry = zip_file.glob("KEN_ALL.CSV").first
  zip_file.extract(entry, "KEN_ALL.CSV")
end

all = CSV.generate do |csv|
  CSV.foreach("KEN_ALL.CSV", encoding: "CP932:UTF-8") do |row|
    zip   = row[2]
    kana  = ConvertWideAndNarrow(NKF.nkf("-Xwh1", row[3] + row[4] + row[5]))
    kanji = ConvertWideAndNarrow(row[6] + row[7] + row[8])
    csv.add_row([zip, kana, kanji])
  end
end

File.open("app/all.csv", "wb") do |f|
  f.write(all)
end

File.delete("ken_all.zip")
File.delete("KEN_ALL.CSV")
