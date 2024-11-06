import { useState, useEffect } from 'react';
import { getCurrentClassroomUser } from '@/api/classrooms';

export function useClassroomUser(classroomId?: number) {
  const [classroomUser, setClassroomUser] = useState<IClassroomUser | null>(null);
  const [error, setError] = useState<Error | null>(null);
  const [loading, setLoading] = useState(true);


  useEffect(() => {
    const fetchClassroomUser = async () => {
      if (classroomId) {
        await getCurrentClassroomUser(classroomId)
        .then((user) => {
          setClassroomUser(user);
          setError(null);
        })
        .catch((err) => {
          setError(err instanceof Error ? err : new Error('Failed to fetch classroom user'));
          setClassroomUser(null);
        })
        .finally(() => {
          setLoading(false);
        });
      } else {
        setClassroomUser(null);
        setError(null);
        setLoading(false);
      }
    };

    fetchClassroomUser();
  }, [classroomId]);

  return { classroomUser, error, loading };
}
