using Microsoft.EntityFrameworkCore;

namespace CoproXpert.Database.Services;

public abstract class BaseService<T> where T : class
{
    private readonly DataContext _context;

    protected BaseService(DataContext context)
    {
        _context = context;
    }

    public virtual IEnumerable<T> GetAll()
    {
        return _context.Set<T>().ToList();
    }

    public virtual T? GetById(int id)
    {
        return _context.Set<T>().Find(id);
    }

    public virtual T Create(T entity)
    {
        _context.Set<T>().Add(entity);
        _context.SaveChanges();
        return entity;
    }

    public virtual bool Update(T entity)
    {
        try
        {
            _context.Entry(entity).State = EntityState.Modified;
            _context.SaveChanges();
            return true;
        }
        catch (DbUpdateConcurrencyException)
        {
            return false;
        }
    }

    public virtual bool Delete(int id)
    {
        var entity = _context.Set<T>().Find(id);
        if (entity == null) return false;
        _context.Set<T>().Remove(entity);
        _context.SaveChanges();
        return true;
    }
}