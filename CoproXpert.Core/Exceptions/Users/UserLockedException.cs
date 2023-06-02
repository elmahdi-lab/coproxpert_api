// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

namespace CoproXpert.Core.Exceptions.Users;

public class UserLockedException : Exception
{
    public UserLockedException()
    {
    }

    public UserLockedException(string message) : base(message)
    {
    }

    public UserLockedException(string message, Exception innerException) : base(message, innerException)
    {
    }
}
