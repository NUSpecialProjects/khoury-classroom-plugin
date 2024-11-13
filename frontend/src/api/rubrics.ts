const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;

export const createRubric = async (
    rubric: IFullRubric
  ): Promise<IFullRubric> => {
    const response = await fetch(
      `${base_url}/rubrics/rubric`,
      {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(rubric),
      }
    );
    if (!response.ok) {
      throw new Error("Network response was not ok");
    }

    const data: IFullRubric = (await response.json() as IFullRubric) 

    return data
  
  };