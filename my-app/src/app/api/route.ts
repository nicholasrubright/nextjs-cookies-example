import { NextResponse } from "next/server";

export async function POST() {
  const apiResponse = await fetch("http://localhost:8080/set", {
    method: "POST",
    cache: "no-cache",
  });

  const data = await apiResponse.json();

  const response = NextResponse.json(data);

  return response;
}
