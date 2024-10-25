import React, { useContext } from "react";
import GenericRolePage from "..";
import LinkGenerator from "@/components/LinkGenerator";
import { SelectedSemesterContext } from "@/contexts/selectedSemester";


const TAListPage: React.FC = () => {
    const { selectedSemester } = useContext(SelectedSemesterContext);
    const role_type = "TA";
    return (
        <>
        <GenericRolePage role_type={role_type} />
            <div>
                <p>Add {role_type}</p>
                <LinkGenerator role_type={role_type} semester={selectedSemester} />
            </div>
        </>
    );
};

export default TAListPage;