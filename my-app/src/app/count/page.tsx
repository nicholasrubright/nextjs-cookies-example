import { cookies } from "next/dist/client/components/headers";

export default async function Page() {
  const response = await getCount();

  let data;

  if (response.ok) {
    data = await response.json();
  } else {
    data = { count: "balls" };
  }

  return (
    <div>
      <h1>Count Page!</h1>
      <p>Count: {data.count}</p>
    </div>
  );
}

async function getCount() {
  return await fetch("http://localhost:3000/api/count", {
    method: "GET",
    cache: "no-cache",
    headers: {
      Cookie: `mysession=${cookies().get("mysession")?.value}`,
    },
  });
}
