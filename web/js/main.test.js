import * as main from "./main";

describe("function buildKeywordArray", () => {
  it("単一", () => {
    const result = main.buildKeywordArray("abc");
    expect(result).toEqual(["abc"]);
  });

  it("スペースあり", () => {
    const result = main.buildKeywordArray("a b   c");
    expect(result).toEqual(["a", "b", "c"]);
  });

  it("スペース、ハイフンあり", () => {
    const result = main.buildKeywordArray("he　l-l　 - 　　o-");
    expect(result).toEqual(["he", "ll", "o"]);
  });

  it("スペース、ハイフン、全角数字あり", () => {
    const result = main.buildKeywordArray("　　 　hel-lo 1２３-45-６７ ８90  　　 ");
    expect(result).toEqual(["hello", "1234567", "890"]);
  });

  it("文字数制限(50文字)", () => {
    const result = main.buildKeywordArray(
      "０１２３４５６７８９　　　　　　　　　　あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほんんんんん"
    );
    expect(result).toEqual(["0123456789", "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほ"]);
  });
});

describe("function getAllIndexes", () => {
  it("該当なし", () => {
    const result = main.getAllIndexes("abc", "z");
    expect(result).toEqual([]);
  });

  it("該当１つ", () => {
    const result = main.getAllIndexes("あいう", "い");
    expect(result).toEqual([1]);
  });

  it("該当２つ", () => {
    const result = main.getAllIndexes("あいういあいうあうい", "うい");
    expect(result).toEqual([2, 8]);
  });

  it("全て該当", () => {
    const result = main.getAllIndexes("あああああ", "あ");
    expect(result).toEqual([0, 1, 2, 3, 4]);
  });
});

describe("function returnMatchArray", () => {
  it("該当なし", () => {
    const result = main.returnMatchArray(["z"], "abcde");
    expect(result).toEqual([false, false, false, false, false]);
  });

  it("該当あり", () => {
    const result = main.returnMatchArray(["a"], "abcaa");
    expect(result).toEqual([true, false, false, true, true]);
  });

  it("該当あり・なし", () => {
    const result = main.returnMatchArray(["a", "z"], "asdf");
    expect(result).toEqual([true, false, false, false]);
  });

  it("複数キーワード該当", () => {
    const result = main.returnMatchArray(["い", "えお"], "あいうえお");
    expect(result).toEqual([false, true, false, true, true]);
  });

  it("キーワード複数、オーバーラップあり", () => {
    const result = main.returnMatchArray(["うえお", "えおか"], "あいうえおかき");
    expect(result).toEqual([false, false, true, true, true, true, false]);
  });
});

describe("function insertString", () => {
  it("先頭に付加", () => {
    const result = main.insertString("abcde", "xyz", 0);
    expect(result).toBe("xyzabcde");
  });

  it("途中に付加", () => {
    const result = main.insertString("abcde", "xyz", 2);
    expect(result).toBe("abxyzcde");
  });

  it("末尾に付加", () => {
    const result = main.insertString("abcde", "xyz", 5);
    expect(result).toBe("abcdexyz");
  });

  it("レンジ外指定", () => {
    const result = main.insertString("abcde", "xyz", 99);
    expect(result).toBe("abcdexyz");
  });
});

describe("function buildSpanEmbeddedHTML", () => {
  it("全一致", () => {
    const result = main.buildSpanEmbeddedHTML("abcde", [true, true, true, true, true]);
    expect(result).toBe('<span class="matched">abcde</span>');
  });

  it("複数の一致", () => {
    const result = main.buildSpanEmbeddedHTML("abcde", [false, true, false, true, false]);
    expect(result).toBe('a<span class="matched">b</span>c<span class="matched">d</span>e');
  });

  it("先頭のみ一致", () => {
    const result = main.buildSpanEmbeddedHTML("abcde", [true, false, false, false, false]);
    expect(result).toBe('<span class="matched">a</span>bcde');
  });

  it("末尾のみ一致", () => {
    const result = main.buildSpanEmbeddedHTML("abcde", [false, false, false, false, true]);
    expect(result).toBe('abcd<span class="matched">e</span>');
  });

  it("先頭に連続した一致", () => {
    const result = main.buildSpanEmbeddedHTML("abcde", [true, true, false, false, false]);
    expect(result).toBe('<span class="matched">ab</span>cde');
  });

  it("末尾に連続した一致", () => {
    const result = main.buildSpanEmbeddedHTML("abcde", [false, false, false, true, true]);
    expect(result).toBe('abc<span class="matched">de</span>');
  });

  it("途中に連続した一致", () => {
    const result = main.buildSpanEmbeddedHTML("abcde", [false, true, true, false, false]);
    expect(result).toBe('a<span class="matched">bc</span>de');
  });

  it("一致なし", () => {
    const result = main.buildSpanEmbeddedHTML("abcde", [false, false, false, false, false]);
    expect(result).toBe("abcde");
  });
});

describe("function buildResultHTML", () => {
  it("かなと漢字のマッチハイライト、漢字欄における数字は全角化される", () => {
    const result = main.buildResultHTML({ zipcode: "0123456", kana: "とうきょう", kanji: "東京123" }, "と　京 1-");
    let expected = "";
    expected += `<div class="address">`;
    expected += `<div class="zipcode">`;
    expected += `<a href=https://www.post.japanpost.jp/cgi-zip/zipcode.php?zip=0123456 target=_blank>`;
    expected += `〒012-3456</a>`;
    expected += `</div>`;
    expected += `<div class="kana-kanji"><div class="kana"><span class="matched">と</span>うきょう</div>`;
    expected += `<div class="kanji">東<span class="matched">京１</span>２３</div>`;
    expected += `</div></div>`;
    expect(result).toBe(expected);
  });
});
