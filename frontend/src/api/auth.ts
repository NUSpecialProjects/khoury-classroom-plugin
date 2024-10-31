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
