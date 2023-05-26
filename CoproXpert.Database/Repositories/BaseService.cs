// Copyright (c) COPRO XPERT - IT HUMANS  All Rights Reserved.

using CoproXpert.Database;
using Microsoft.EntityFrameworkCore;

namespace CoProXpert.Database.Repositories;

public abstract class BaseService<T> where T : class
{
    protected readonly DataContext Context;

    protected BaseService(DataContext context)
    {
        Context = context;
    }

    public virtual IEnumerable<T> GetAll()
    {
        return Context.Set<T>().ToList();
    }

    public virtual T? GetById(int id)
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

    public virtual bool Delete(int id)
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
