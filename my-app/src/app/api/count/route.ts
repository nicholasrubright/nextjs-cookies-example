import { cookies } from "next/dist/client/components/headers";
import { NextRequest, NextResponse } from "next/server";

export async function GET(request: NextRequest) {
  const apiResponse = await fetch("http://localhost:8080/get", {
    method: "GET",
    cache: "no-cache",
    headers: {
      Cookie: `mysession=${cookies().get("mysession")?.value}`,
    },
  });

  const apiData = await apiResponse.json();

  const response = NextResponse.json(apiData);

  return response;
}
