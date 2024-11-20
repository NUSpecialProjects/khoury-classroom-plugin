const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;

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

export const sendCode = async (code: string): Promise<void> => {
  const response = await fetch(`${base_url}/login`, {
    method: "POST",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ code }),
  });

  if (!response.ok) {
    const errorData = await response.json();
    throw new Error(errorData.message || "An error occurred during login");
  }

  return;
};

export const logout = async (): Promise<void> => {
  const response = await fetch(`${base_url}/logout`, {
    method: "POST",
    credentials: "include",
  });

  if (!response.ok) {
    const errorData = await response.json();
    throw new Error(errorData.message || "An error occurred during logout");
  }

  return;
};
