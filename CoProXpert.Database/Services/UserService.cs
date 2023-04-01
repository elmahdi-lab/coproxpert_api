using CoproXpert.Database.Models;

namespace CoproXpert.Database.Services;

public class UserService : BaseService<User>
{
    public UserService(DataContext context) : base(context)
    {
    }
}