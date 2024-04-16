/**
 * @param {Blob} blob
 * @returns {Promise<string>}
 */
async function blobToB64(blob) {
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
    var blob = await fetch(url).then(r => r.blob());
    var b64data = await blobToB64(blob);
    return {
        "contentType": blob.type,
        "b64data": b64data.split(",")[1],
    };
}