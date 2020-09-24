"use strict";

import $ from "jquery";

$(function () {
  let tid = 0;
  let delay = 100;
  let prevWord = "";

  function delayed_search(force) {
    clearTimeout(tid);
    let keyword = $(".search-form").val().trim();
    if (keyword === "") {
      $(".search-results").empty();
      prevWord = "";
      return;
    }
    let prefCode = 0;
    if ($(".pref-limit").prop("checked")) {
      prefCode = $(".pref-select").val();
    }
    if (prevWord !== keyword || force) {
      tid = setTimeout(search, delay, keyword, prefCode);
    }
  }

  $(".search-form").on("input", function () {
    delayed_search(false);
  });

  $(".pref-limit, .pref-select").on("change", function () {
    $(".pref-select").prop("disabled", !$(".pref-limit").prop("checked"));
    delayed_search(true);
  });

  function search(keyword, prefCode) {
    prevWord = keyword;
    $.get({
      url: "/search",
      data: { keyword: keyword, pref: prefCode },
      dataType: "json",
      success: function (data) {
        if (data && data.length > 0) {
          let html = "";
          for (let i = 0, len = data.length; i < len; i++) {
            html += buildResultHTML(data[i], keyword);
          }
          $(".search-results").html(html);
        } else {
          $(".search-results").text("検索結果なし");
        }
      },
      error: function () {
        $(".search-results").empty();
      },
    });
  }
});

export function buildKeywordArray(raw) {
  if (raw === undefined) {
    return [];
  }
  let keyword = "";
  keyword = String(raw).substring(0, 50);
  keyword = keyword.replace(/[-ｰ−－]+/g, "");
  keyword = keyword.replace(/[\s　]+/g, " ");
  keyword = keyword.replace(/[０-９]/g, function (s) {
    return String.fromCharCode(s.charCodeAt(0) - 65248);
  });
  keyword = keyword.trim();
  if (keyword.length === 0) {
    return [];
  }
  return keyword.split(" ");
}

export function getAllIndexes(str, substr) {
  let indexes = [];
  let i = -1;
  while ((i = str.indexOf(substr, i + 1)) != -1) {
    indexes.push(i);
  }
  return indexes;
}

export function returnMatchArray(keywordArray, str) {
  let matchArray = new Array(str.length);
  for (let i = 0, len = matchArray.length; i < len; i++) {
    matchArray[i] = false;
  }
  let examineStr = "";
  let indexArray = [];
  for (let i = 0, len = keywordArray.length; i < len; i++) {
    examineStr = keywordArray[i];
    indexArray = getAllIndexes(str, examineStr);
    for (let k = 0, len = indexArray.length; k < len; k++) {
      for (let m = indexArray[k]; m < indexArray[k] + examineStr.length; m++) {
        matchArray[m] = true;
      }
    }
  }
  return matchArray;
}

export function insertString(target, insstr, index) {
  return target.slice(0, index) + insstr + target.slice(index);
}

export function buildSpanEmbeddedHTML(str, matchArray) {
  let flag = false;
  let html = str;
  for (let i = matchArray.length - 1; i >= 0; i--) {
    if (!flag && matchArray[i]) {
      html = insertString(html, "</span>", i + 1);
      flag = true;
    }
    if (flag && !matchArray[i]) {
      html = insertString(html, '<span class="matched">', i + 1);
      flag = false;
    }
  }
  if (flag) {
    html = insertString(html, '<span class="matched">', 0);
  }
  return html;
}

export function buildResultHTML(data, keyword) {
  let result = "";

  let keywordArray = buildKeywordArray(keyword);

  let kanaMatchArray = returnMatchArray(keywordArray, data.kana);
  let kanjiMatchArray = returnMatchArray(keywordArray, data.kanji);

  let kanahtml = buildSpanEmbeddedHTML(data.kana, kanaMatchArray);
  let kanjihtml = buildSpanEmbeddedHTML(data.kanji, kanjiMatchArray);

  kanahtml = kanahtml.replace(/[0-9]/g, function (s) {
    return String.fromCharCode(s.charCodeAt(0) + 65248);
  });

  kanjihtml = kanjihtml.replace(/[0-9]/g, function (s) {
    return String.fromCharCode(s.charCodeAt(0) + 65248);
  });

  result += `<div class="address">`;
  result += `<div class="zipcode">`;
  result += `<a href=https://www.post.japanpost.jp/cgi-zip/zipcode.php?zip=${data.zipcode} target=_blank>`;
  result += `〒${data.zipcode.slice(0, 3)}-${data.zipcode.slice(3)}</a>`;
  result += `</div>`;
  result += `<div class="kana-kanji"><div class="kana">${kanahtml}</div>`;
  result += `<div class="kanji">${kanjihtml}</div>`;
  result += `</div></div>`;
  return result;
}
