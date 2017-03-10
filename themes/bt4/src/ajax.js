export const get = (url, success) => {
    fetch(
      `${process.env.REACT_APP_BACKEND}${url}`,
      {
        method: 'GET',
        credentials: 'include',
      }
    )
        .then(response => (response.json()))
        .then(success);
}
