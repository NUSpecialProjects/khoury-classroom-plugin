const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;

export const sendLoginRequest = async (code: string): Promise<Response> => {
  const response = await fetch(`${base_url}/login`, {
    method: "POST",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ code }),
  });
  return response;
};

export const getCallbackURL = async (): Promise<string> => {
    const response = await fetch(`${base_url}/callback`, {
        method: "GET",
        credentials: "include",
        headers: {
        "Content-Type": "application/json",
        },
    });
    const data = await response.json();
    return data.url;
};
