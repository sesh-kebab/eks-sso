export const LIST_NAMESPACES = 'namespaces:list';

export function getNamespaces() {
  return dispatch => {
    getData('/namespaces', ).then(o => {
      return o.json();
    })
      .then(j => {
        dispatch(send(j))
      })
      .catch(e => {
        console.log(e.message)
        // dispatch(addError(e.message))
      })
  }
}

const send = (namespaces) => {
  return {
    type: LIST_NAMESPACES,
    payload: {
      namespaces,
    }
  }
}

const getData = (url) => {
  return fetch(url, {
    headers: {
      'Content-Type': 'application/json; charset=utf-8',
    },
    cache: 'no-cache',
    method: 'GET',
    credentials: 'same-origin',
  })
}