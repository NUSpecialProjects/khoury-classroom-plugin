// ... imports remain the same ...

import EmptyDataBanner from "@/components/EmptyDataBanner";
import useUrlParameter from "@/hooks/useUrlParameter";
import { useEffect, useState } from "react";
import ClipLoader from "react-spinners/ClipLoader";

export interface TokenHandlerConfig<T extends ITokenUseResponse> {
        useTokenFunction: (token: string) => Promise<T>;
        successCallback: (response: T) => void;
        loadingMessage?: string;
        successMessage?: (response: T) => string;
  }
  
  const TokenApplyPage = <T extends ITokenUseResponse>({
    useTokenFunction = async () => {throw new Error("useTokenFunction not implemented")},
    successCallback = () => {throw new Error("successCallback not implemented")},
    loadingMessage = "Applying token...",
    successMessage = () => "Success! Redirecting...",
}: TokenHandlerConfig<T>) => {
    const inputToken = useUrlParameter("token");
    const [message, setMessage] = useState<string>("Loading...");
    const [loading, setLoading] = useState<boolean>(true);
    const [success, setSuccess] = useState<boolean>(false);
  
    useEffect(() => {
      if (inputToken && !success) {
        console.log("Input token:", inputToken);
        handleUseToken();
      }
    }, [inputToken]);
  
    const handleUseToken = async () => {
        console.log("Handling use token");
      setLoading(true);
      setMessage(loadingMessage);
      await useTokenFunction(inputToken)
        .then((response) => {
          console.log("Successfully used token");
          setMessage(successMessage(response));
          setSuccess(true);
          successCallback(response);
        })
        .catch((error) => {
          console.log("Error using token: " + error);
          setMessage("Error using token: " + error);
        })
        .finally(() => {
          setLoading(false);
          console.log("Loading set to false");
        });
    };
  
    return (
        <EmptyDataBanner>
            <ClipLoader size={50} color={"#123abc"} loading={loading} />
            <p>{message}</p>
        </EmptyDataBanner>
    );
  };
  
  export default TokenApplyPage;