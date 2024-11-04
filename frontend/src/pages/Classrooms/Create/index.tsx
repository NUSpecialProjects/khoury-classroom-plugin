import React, { useContext, useEffect, useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import { postClassroom } from "@/api/classrooms";
import "./styles.css";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import useUrlParameter from "@/hooks/useUrlParameter";
import { getOrganizationDetails } from "@/api/organizations";

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
    <div className="ClassroomCreation">
      <h1>Create a New Classroom</h1>
      {loading ? (
        <p>Loading...</p>
      ) : (
        <form onSubmit={handleSubmit}>
          <div className="form-group">
            <label htmlFor="name">Classroom Name</label>
            <input
              type="text"
              id="name"
              value={name}
              onChange={(e) => setName(e.target.value)}
              required
            />
          </div>
          <div className="form-group">
            <label htmlFor="organization">Organization</label>
            <input
              type="text"
              id="organization"
              value={organization ? organization.login : ""}
              readOnly
              required
            />
            {!organization && (
              <p className="error">
                Organization not provided. <Link to="/app/organization/select">Click here to select an organization</Link>.
              </p>
            )}
          </div>
          {error && <p className="error">{error}</p>}
          <button type="submit" className="btn btn-primary">
            Create Classroom
          </button>
        </form>
      )}
    </div>
  );
};

export default ClassroomCreation;