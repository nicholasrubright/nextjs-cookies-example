import { NextResponse } from "next/server";

export async function POST() {
  const apiResponse = await fetch("http://localhost:8080/balls/set", {
    method: "POST",
    cache: "no-cache",
  });

  const sessionCookie = apiResponse.headers.getSetCookie()[0];
  let response = NextResponse.json(null);

  if (sessionCookie) {
    response.headers.set("Set-Cookie", sessionCookie);
  }

  return response;
}
