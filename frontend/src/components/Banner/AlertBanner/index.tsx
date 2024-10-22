import React, { useEffect, useState } from "react";
import {
  activateSemester,
  deactivateSemester,
  getOrgSemesters,
} from "@/api/semesters";
import "./styles.css";
interface AlertBannerProps {
  semester: ISemester;
  onActivate: (newSemester: ISemester) => void;
}

enum SemesterError {
  API_ERROR = "Failed to modify class. Please try again.",
  ALREADY_ACTIVE = "A class is already active. Please deactivate it first.",
  MULTIPLE_ACTIVE = "Multiple classes are active. Please deactivate all but one.",
  NOT_ACTIVE = "This class is not active. Activate it to use this class.",
}

const AlertBanner: React.FC<AlertBannerProps> = ({ semester, onActivate }) => {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [activeSemester, setActiveSemester] = useState<ISemester | null>(null);

  useEffect(() => {
    const checkErrors = async () => {
      const orgSemesters: ISemester[] = (await getOrgSemesters(semester.org_id))
        .semesters;
      const activeSemesters = orgSemesters.filter((s: ISemester) => s.active);
      const otherActiveSemester = activeSemesters.find(
        (s: ISemester) => s.classroom_id !== semester.classroom_id
      );
      if (activeSemesters.length > 1) {
        setError(SemesterError.MULTIPLE_ACTIVE);
      } else if (otherActiveSemester) {
        setActiveSemester(otherActiveSemester);
        setError(SemesterError.ALREADY_ACTIVE);
      } else if (!semester.active) {
        setError(SemesterError.NOT_ACTIVE);
      } else if (semester.active && error !== SemesterError.API_ERROR) {
        setError(null);
      }
    };
    void checkErrors();
  }, [semester]);

  const handleActivate = async () => {
    setLoading(true);
    setError(null);
    try {
      const newSemester = await activateSemester(
        semester.org_id,
        semester.classroom_id
      );
      onActivate(newSemester);
    } catch (err) {
      console.log(err);
      setError(SemesterError.API_ERROR);
    } finally {
      setLoading(false);
    }
  };

  const handleDeactivate = async () => {
    setLoading(true);
    setError(null);
    try {
      const newSemester = await deactivateSemester(
        semester.org_id,
        semester.classroom_id
      );
      onActivate(newSemester);
    } catch (err) {
      console.log(err);
      setError(SemesterError.API_ERROR);
    } finally {
      setLoading(false);
    }
  };

  const handleClick = () => {
    handleActivate().catch((err: unknown) => {
      console.error("Error in handleActivate:", err);
    });
  };

  const handleNavigateToActiveSemester = () => {
    if (
      activeSemester &&
      activeSemester.org_id &&
      activeSemester.classroom_id
    ) {
      onActivate(activeSemester);
    }
  };

  return (
    error && (
      <div className="inactive-class-banner">
        {error && <p className="error">{error}</p>}
        {error === SemesterError.ALREADY_ACTIVE && activeSemester && (
          <button onClick={handleNavigateToActiveSemester}>
            Go to Active Semester
          </button>
        )}
        {semester.active && error === SemesterError.MULTIPLE_ACTIVE && (
          <button onClick={handleDeactivate}>Deactivate this semester</button>
        )}
        {(error === SemesterError.API_ERROR ||
          error === SemesterError.NOT_ACTIVE ||
          (!semester.active && error === null)) && (
          <button onClick={handleClick} disabled={loading}>
            {loading ? "Activating..." : "Activate Class"}
          </button>
        )}
      </div>
    )
  );
};

export default AlertBanner;
