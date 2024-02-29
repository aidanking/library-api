const baseUrl = 'http://localhost:8081/api';

async function makeRequest(method = 'GET', url, body = '') {
  const headers = {
    accepts: 'application/json',
    'Content-Type': 'application/json',
  };
  try {
    const resp = await fetch(url, {
      method,
      headers,
      body: body ? JSON.stringify(body) : null,
    });
    const json = await resp.json();
    return json;
  } catch (err) {
    console.log(err);
  }
}

(async function main() {
  let id;
  const createdAuthor = await makeRequest('POST', `${baseUrl}/authors`, {
    firstName: 'Aidan',
    lastName: 'King 2',
    country: 'IE',
  });

  id = createdAuthor.id;
  console.log(`created id ${id}`);
  console.log('created author');
  console.log(createdAuthor);

  const author = await makeRequest('GET', `${baseUrl}/authors/${id}`);
  console.log('author');
  console.log(author);

  const updatedAuthor = await makeRequest('PUT', `${baseUrl}/authors/${id}`, {
    id: 8,
    firstName: 'Test',
    lastName: 'King',
    country: 'IE689',
  });
  console.log('updated author');
  console.log(updatedAuthor);

  const deletedAuthor = await makeRequest('DELETE', `${baseUrl}/authors/${id}`);
  console.log('deleted author');
  console.log(deletedAuthor);

  const allAuthors = await makeRequest('GET', `${baseUrl}/authors`);
  console.log('all authors');
  console.log(allAuthors);

  const requiredCountryNameOnCreate = await makeRequest(
    'POST',
    `${baseUrl}/authors`,
    {
      firstName: 'Aidan',
      lastName: 'King 2',
    }
  );
  console.log('country name required on create');
  console.log(requiredCountryNameOnCreate);

  const requiredFirstNameOnUpdate = await makeRequest(
    'PUT',
    `${baseUrl}/authors/${id}`,
    {
      id: 8,
      lastName: 'King',
      country: 'IE689',
    }
  );
  console.log('first name is required on update');
  console.log(requiredFirstNameOnUpdate);
})();
