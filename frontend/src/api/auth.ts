const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;

export const getCurrentUser = async (): Promise<boolean> => {
  const result = await fetch(`${base_url}/github/user`, {
    method: "GET",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
  });

  return result.ok;
};
