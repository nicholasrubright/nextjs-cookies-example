export default async function Page() {
  const data = await setCount();
  return (
    <div>
      <h1>Hello world!</h1>
      <p>Response: {data.count}</p>
    </div>
  );
}

async function setCount() {
  const response = await fetch("http://localhost:3000/api", {
    method: "POST",
    cache: "no-cache",
    credentials: "include",
  });

  return response.json();
}
