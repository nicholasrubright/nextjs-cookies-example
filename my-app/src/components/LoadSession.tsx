"use client";

import { useEffect } from "react";

export default function LoadSession(props: LoadSessionProps) {
  const { hasCookie } = props;

  useEffect(() => {
    const setSession = async () => {
      await fetch("http://localhost:3000/api", {
        method: "POST",
      });
    };

    if (!hasCookie) {
      setSession();
    }
  }, [hasCookie]);

  return null;
}

interface LoadSessionProps {
  hasCookie: boolean;
}
