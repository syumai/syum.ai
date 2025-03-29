async function main() {
  const avatarEl = document.getElementById("avatar");
  const shareButton = document.getElementById("shareButton");
  const reloadButton = document.getElementById("reloadButton");

  let avatarDataUrl = "";
  let colorCode = avatarEl.dataset.initialColorCode;

  const setAvatarImage = async () => {
    try {
      const url = "/image/random";
      const result = await fetch(url);
      if (avatarDataUrl !== "") {
        URL.revokeObjectURL(avatarDataUrl);
      }
      avatarDataUrl = URL.createObjectURL(await result.blob());
      colorCode = URL.parse(result.url).searchParams.get("code");
    } catch (e) {
      console.error(e);
      return;
    }
    avatarEl.src = avatarDataUrl;
    // support only 200px image for shuffle button.
    avatarEl.srcset = "";
  };

  avatarEl.addEventListener("load", (e) => {
    e.target.classList.remove("is-invisible");
  });

  reloadButton.addEventListener("click", async (e) => {
    await setAvatarImage();
  });

  if (window.navigator.share) {
    shareButton.addEventListener("click", (e) => {
      window.navigator.share({
        title: "syum.ai",
        text: "ランダムシュウマイで遊んでみよう! #わたしのシュウマイ",
        url: `https://syum.ai?colorCode=${colorCode}`,
      });
    });
  }
}

window.addEventListener("DOMContentLoaded", main);
