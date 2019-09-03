async function main() {
  const avatarEl = document.getElementById('avatar');
  const sharedUrlEl = document.getElementById('sharedUrl');
  const reloadButton = document.getElementById('reloadButton');

  let avatarDataUrl = '';

  const setAvatarImage = async () => {
    let sharedUrl;
    try {
      const result = await fetch('/image/random');
      if (avatarDataUrl !== '') {
        URL.revokeObjectURL(avatarDataUrl);
      }
      avatarDataUrl = URL.createObjectURL(await result.blob());
      sharedUrl = result.url;
    } catch (e) {
      console.error(e);
      return;
    }
    avatarEl.src = avatarDataUrl;
    sharedUrlEl.href = sharedUrl;
    sharedUrlEl.textContent = sharedUrl;
  };

  reloadButton.addEventListener('click', async (e) => {
    e.preventDefault();
    await setAvatarImage();
  });

  if (window.navigator.share) {
    sharedUrlEl.addEventListener('click', (e) => {
      e.preventDefault();
      window.navigator.share({
        title: '#わたしのシュウマイ',
        text: '#わたしのシュウマイ',
        url: e.target.textContent
      });
    })
  }

  await setAvatarImage();
}

window.addEventListener('DOMContentLoaded', main);
