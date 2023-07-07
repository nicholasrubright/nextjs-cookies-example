export default async function Page() {
  const response = await setCount();

  const data = await response.json();

  return (
    <div>
      <h1>Hello world!</h1>
      <p>Response: {data.count}</p>
    </div>
  );
}

async function setCount() {
  return await fetch("http://localhost:3000/api", {
    method: "POST",
    cache: "no-cache",
    credentials: "include",
  });
}
