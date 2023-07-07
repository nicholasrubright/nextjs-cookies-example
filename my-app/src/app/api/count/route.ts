import { cookies } from "next/headers";
import { NextResponse } from "next/server";

export async function GET() {
  let data = { token: "balls" };

  const sessionCookie = cookies().get("mysession")?.value;

  if (sessionCookie) {
    const apiResponse = await fetch("http://localhost:8080/api/get", {
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
