// https://ww7.mangakakalot.tv/chapter/manga-lm962995/chapter-1
// thank god for these guys and their normal-ass cdn

const inputUrls = [
  "https://cm.blazefast.co/35/7b/357bdfddfb5a7b521d74d1b623699fdc.jpg",
  "https://cm.blazefast.co/30/20/3020dcf00566ddea84dae2f5524e23b1.jpg",
  "https://cm.blazefast.co/a6/9f/a69f1529b2d29db6d343dc8397f7f937.jpg",
  "https://cm.blazefast.co/6f/ad/6fadf71580d2800984163e0b090c2a52.jpg",
  "https://cm.blazefast.co/eb/75/eb750d15d03392f3d23c4cb05535cbb8.jpg",
  "https://cm.blazefast.co/4f/71/4f71f154ab0192690d5157a68065106a.jpg",
  "https://cm.blazefast.co/c9/24/c9247e97de0d8366a4fa0b90a071452f.jpg",
  "https://cm.blazefast.co/ea/b0/eab070fe517702ae3d6de7dec2bda2c5.jpg",
  "https://cm.blazefast.co/6b/e6/6be6f32b407e9e14e56fed482648da61.jpg",
  "https://cm.blazefast.co/61/6a/616abaeede4981ce7f720dc5f9508e14.jpg",
  "https://cm.blazefast.co/77/49/77495aca138ccd41b2c384cc096c78b5.jpg",
  "https://cm.blazefast.co/dd/5a/dd5a8ec12424f51fcaa4b58bd6d7d6db.jpg",
  "https://cm.blazefast.co/c1/19/c1194948abe644384f52ea4e9f0dc226.jpg",
  "https://cm.blazefast.co/53/63/5363412e5ee4055cd048e488c567dc0f.jpg",
  "https://cm.blazefast.co/c3/18/c318788cdc16f660671438b5e307cf46.jpg",
  "https://cm.blazefast.co/5f/c6/5fc6c10130481dc5c6c36c989cfc2d37.jpg",
  "https://cm.blazefast.co/8d/58/8d58ab3e456ac6796e88a870d81c19f2.jpg",
  "https://cm.blazefast.co/ef/67/ef6738c8f9caa0b8e0f9cefbc7385bec.jpg",
  "https://cm.blazefast.co/51/5b/515b39b4b5598eb452849fc092a6a57b.jpg",
  "https://cm.blazefast.co/cc/e3/cce3e6fc07eb80306b39887072628711.jpg",
  "https://cm.blazefast.co/dc/a2/dca22f7616aacc943101a08803a90ac9.jpg",
  "https://cm.blazefast.co/31/9b/319b8c6aa376a17a26070f9daa6e273b.jpg",
  "https://cm.blazefast.co/c2/29/c229fd43c1cce0a7493127ab27efb958.jpg",
  "https://cm.blazefast.co/ed/e2/ede2956b97d7e10eb33ac3a4c07c2795.jpg",
];

await Promise.all(
  (await Promise.all(inputUrls.map(fetch))).map((res, i) => {
    if (res.ok) {
      Bun.write(String(i + 1).padStart(3, "0") + ".jpg", res);
    } else {
      console.log(res);
    }
  }),
);
