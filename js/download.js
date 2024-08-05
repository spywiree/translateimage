/**
 * @param {Blob} blob
 * @returns {Promise<string>}
 */
async function blobToBase64(blob) {
  return new Promise((resolve) => {
    const reader = new FileReader();
    reader.onloadend = () => resolve(reader.result);
    reader.readAsDataURL(blob);
  });
}

/**
 * @param {string} url
 */
async function download(url) {
  const blob = await fetch(url).then((r) => r.blob());
  const b64data = await blobToBase64(blob);
  return {
    contentType: blob.type,
    b64data: b64data.split(",")[1],
  };
}
