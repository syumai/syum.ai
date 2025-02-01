async function main() {
  const avatarEl = document.getElementById("avatar");
  const shareButton = document.getElementById("shareButton");
  const reloadButton = document.getElementById("reloadButton");

  // URLからcolorCodeパラメータを取得する
  let initialColorCode = new URLSearchParams(window.location.search).get(
    "colorCode"
  );
  let avatarDataUrl = "";
  let sharedUrl;

  const setAvatarImage = async () => {
    try {
      let url = "/image/random";
      if (initialColorCode) {
        url = `/image?code=${initialColorCode}`;
        initialColorCode = "";
      }
      const result = await fetch(url);
      if (avatarDataUrl !== "") {
        URL.revokeObjectURL(avatarDataUrl);
      }
      avatarDataUrl = URL.createObjectURL(await result.blob());
      sharedUrl = result.url;
    } catch (e) {
      console.error(e);
      return;
    }
    avatarEl.src = avatarDataUrl;
  };

  avatarEl.addEventListener("load", (e) => {
    e.target.classList.remove("is-invisible");
  });

  reloadButton.addEventListener("click", async (e) => {
    await setAvatarImage();
  });

  if (window.navigator.share) {
    shareButton.addEventListener("click", (e) => {
      colorCode = URL.parse(sharedUrl).searchParams.get("code");
      window.navigator.share({
        title: "syum.ai",
        text: "#わたしのシュウマイ",
        url: `https://syum.ai?colorCode=${colorCode}`,
      });
    });
  }

  await setAvatarImage();
}

window.addEventListener("DOMContentLoaded", main);
