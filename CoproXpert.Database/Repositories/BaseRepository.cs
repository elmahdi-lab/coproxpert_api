// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using Microsoft.EntityFrameworkCore;

namespace CoproXpert.Database.Repositories;

public abstract class BaseRepository<T> where T : class
{
    protected BaseRepository(DataContext context)
    {
        Context = context;
    }

    private protected DataContext Context { get; }

    public virtual IEnumerable<T> GetAll()
    {
        return Context.Set<T>().ToList();
    }

    public virtual T? GetById(Guid id)
    {
        return Context.Set<T>().Find(id);
    }

    public virtual T Create(T entity)
    {
        Context.Set<T>().Add(entity);
        Context.SaveChanges();
        return entity;
    }

    public virtual bool Update(T entity)
    {
        try
        {
            Context.Entry(entity).State = EntityState.Modified;
            Context.SaveChanges();
            return true;
        }
        catch (DbUpdateConcurrencyException)
        {
            return false;
        }
    }

    public virtual bool Delete(Guid id)
    {
        var entity = Context.Set<T>().Find(id);
        if (entity == null)
        {
            return false;
        }

        Context.Set<T>().Remove(entity);
        Context.SaveChanges();
        return true;
    }
}
