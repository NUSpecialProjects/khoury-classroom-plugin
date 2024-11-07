import React, { useContext, useEffect, useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import { postClassroom } from "@/api/classrooms";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { getOrganizationDetails } from "@/api/organizations";
import useUrlParameter from "@/hooks/useUrlParameter";
import Panel from "@/components/Panel";
import Button from "@/components/Button";
import Input from "@/components/Input";

import "./styles.css";

const ClassroomCreation: React.FC = () => {
  const [name, setName] = useState("");
  const [organization, setOrganization] = useState<IOrganization | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const orgID = useUrlParameter("org_id");
  const { setSelectedClassroom } = useContext(SelectedClassroomContext);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchOrganizationDetails = async () => {
      if (orgID) {
        setLoading(true);
        await getOrganizationDetails(orgID)
          .then((org) => {
            setOrganization(org);
          })
          .catch((error) => {
            console.error("Error fetching organization details:", error);
            setError("Failed to fetch organization details. Please try again.");
          })
          .finally(() => {
            setLoading(false);
          });
      }
    };

    fetchOrganizationDetails();
  }, [orgID, navigate]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!name || !organization) {
      setError("Please fill in all fields.");
      return;
    }
    setLoading(true);
    await postClassroom({
      name: name,
      org_id: organization.id,
      org_name: organization.login,
    })
      .then((createdClassroom) => {
        setSelectedClassroom(createdClassroom);
        console.log("Created classroom:", createdClassroom);
        navigate("/app/dashboard");
      })
      .catch((error) => {
        console.error("Error creating classroom:", error);
        setError("Failed to create classroom. Please try again.");
      })
      .finally(() => {
        setLoading(false);
      });
  };

  return (
    <Panel title="New Classroom" logo={true}>
      <div className="ClassroomCreation">
        {loading ? (
          <p>Loading...</p>
        ) : (
          <form onSubmit={handleSubmit}>
            <Input
              label="Organization name"
              name="organization"
              required
              readOnly
              value={organization ? organization.login : ""}
            />

            <Input
              label="Classroom name"
              name="classroom-name"
              placeholder="Enter a name for your classroom..."
              required
              value={name}
              onChange={(e) => setName(e.target.value)}
            />
            {error && <p className="error">{error}</p>}
            {!organization && (
              <p className="error">
                <Link to="/app/organization/select">
                  Click here to select an organization
                </Link>
                .
              </p>
            )}
            <div className="ClassroomCreation__buttonWrapper">
              <Button type="submit" variant="primary">Create Classroom</Button>
              <Button variant="secondary" onClick={() => navigate("/app/organization/select")}>Select a different organization</Button>
            </div>
          </form>
        )}
      </div>
    </Panel>
  );
};

export default ClassroomCreation;
