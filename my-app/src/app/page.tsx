import LoadSession from "@/components/LoadSession";
import { cookies } from "next/headers";

export default async function Page() {
  return (
    <div>
      <LoadSession hasCookie={cookies().has("mysession")} />
      <h1>Hello world!</h1>
    </div>
  );
}
