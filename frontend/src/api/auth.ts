const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;

export const getCurrentUser = async (): Promise<boolean> => {
  try {
    const result = await fetch(`${base_url}/user`, {
      method: "GET",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    });

    return result.ok;
  } catch (error) {
    console.log("Error fetching current user:", error);
    return false;
  }
  // return Promise.resolve(false);
};

export const sendCode = async (code: string): Promise<Response> => {
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

export const logout = async (): Promise<Response> => {
  const response = await fetch(`${base_url}/logout`, {
    method: "POST",
    credentials: "include",
  });
  return response;
};
