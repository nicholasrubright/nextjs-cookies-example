export default async function Page() {
  const data = await getCount();

  return (
    <div>
      <h1>Count Page!</h1>
      <p>Count: {data.count}</p>
    </div>
  );
}

async function getCount() {
  const response = await fetch("http://localhost:8080/get", {
    method: "GET",
    cache: "no-cache",
  });

  return response.json();
}
