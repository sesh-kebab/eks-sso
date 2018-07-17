const postData = (url, data) => (
    fetch(url, {
        body: JSON.stringify(data),
        headers: {
            'Content-Type': 'application/json; charset=utf-8',
        },
        cache: 'no-cache',
        method: 'POST',
     credentials: 'same-origin',
    })
)

export default { postData };