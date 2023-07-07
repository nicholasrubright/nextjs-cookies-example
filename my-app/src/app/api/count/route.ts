import { cookies } from "next/dist/client/components/headers";
import { NextResponse } from "next/server";

export async function GET() {
  let data = { token: "balls" };

  if (cookies().has("mysession")) {
    const apiResponse = await fetch("http://localhost:8080/get", {
      method: "GET",
      cache: "no-cache",
      headers: {
        Cookie: `mysession=${cookies().get("mysession")?.value}`,
      },
    });

    data = await apiResponse.json();
  }

  const response = NextResponse.json(data);

  return response;
}
